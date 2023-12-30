import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.song import SongSchema
from models.http_exceptions import *



songs_url = "http://localhost:8084/songs"  # URL de l'API users (golang)



def create_song(song_register):
    print(song_register)
    # on récupère le schéma chanson pour la requête vers l'API users
    song_schema = SongSchema().loads(json.dumps(song_register), unknown=EXCLUDE)
    print(song_schema)
    # on crée la chonson côté API users
    response = requests.request(method="POST", url=songs_url, json=song_schema)
     
    
    print(response.json())

    if response.status_code != 201:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

def get_song(id):
    response = requests.request(method="GET", url=songs_url+"/"+id)
    return response.json(), response.status_code


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


def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+"/"+id)
    return  response.status_code

def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code

