import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.song import SongSchema
from models.http_exceptions import *
import services.ratings as ratings_service


songs_url = "http://localhost:8083/songs"  # URL de l'API users (golang)


# create song
def create_song(song_register):
    print(song_register)
    # on récupère le schéma chanson pour la requête vers l'API users
    song_schema = SongSchema().loads(json.dumps(song_register), unknown=EXCLUDE)
    print(song_schema)
    # on crée la chonson côté API users
    response = requests.request(method="POST", url=songs_url, json=song_schema)
     
    
  

    if response.status_code != 201:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

# get song by id
def get_song(id):
    response = requests.request(method="GET", url=songs_url+"/"+id)
    return response.json(), response.status_code

# update song 
def update_song(id, song_update):


    # s'il y a quelque chose à changer côté API 
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    print(song_schema)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=songs_url+"/"+id, json=song_schema)
        print(response.status_code)
        if response.status_code != 200:
            return response.json(), response.status_code

    return response.json(), response.status_code

# delete song
def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+"/"+id)
    if response.status_code == 204:
    # Pas de contenu, retourner une réponse vide ou un message approprié
        return "song supprimée avec succès", 204
    else:
    # Gérer la réponse avec du contenu
        return response.json(), response.status_code    

# get all songs
def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code

# get ratings by song id
def get_ratings(id):
    
    ratings, status_code =  ratings_service.get_ratings()
    if status_code != 200:
        return ratings, status_code
    
    songsRatings = []
    for rating in ratings:
        if rating["music_id"] == id:
            songsRatings.append(rating)
    return songsRatings, 200

# create rating for a song
def create_rating(id, rating):
    rating["music_id"] = id
    rating["user_id"] = current_user.id
    
    response = ratings_service.create_rating(rating)
    if response.status_code != 201:
        return response.json(), response.status_code

    return response.json(), response.status_code 



