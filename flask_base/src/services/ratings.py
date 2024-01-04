import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user
import services.songs as songs_service
from schemas.rating import RatingSchema
from models.http_exceptions import *



ratings_url = "http://localhost:8084/ratings"  # URL de l'API ratings (golang)



def create_rating(rating_register):
    print(rating_register)
    # on récupère le schéma rating pour la requête vers l'API ratings
    rating_schema = RatingSchema().loads(json.dumps(rating_register), unknown=EXCLUDE)
    rating_schema["user_id"] = current_user.id

    #check if music_id exists
    songs, status_code = songs_service.get_song(rating_schema["music_id"])
  

    if status_code != 200:
        return songs, status_code
    
    # on crée la chonson côté API ratings
    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
     


    if response.status_code != 201:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

def get_rating(id):
    response = requests.request(method="GET", url=ratings_url+"/"+id)
    return response.json(), response.status_code


def update_rating(id, rating_update):
    
    
    # check if the rating is from the current user
    rating, status_code = get_rating(id)
    if status_code != 200:
        return rating, status_code
    
    if rating["user_id"] != current_user.id:
        raise Forbidden
    
    # s'il y a quelque chose à changer côté API 
    rating_schema = RatingSchema().loads(json.dumps(rating_update), unknown=EXCLUDE)

    response = None
    if not RatingSchema.is_empty(rating_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=ratings_url+"/"+id, json=rating_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    return response.json(), response.status_code


def delete_rating(id):
    # check if the rating is from the current user
    rating, status_code = get_rating(id)
    if status_code != 200:
        return rating, status_code
    
    if rating["user_id"] != current_user.id:
        raise Forbidden
    
    response = requests.request(method="DELETE", url=ratings_url+"/"+id)
    return response.json(), response.status_code

def get_ratings():
    response = requests.request(method="GET", url=ratings_url)
    return response.json(), response.status_code
