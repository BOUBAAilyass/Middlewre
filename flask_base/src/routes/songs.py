import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from models.http_exceptions import *
from schemas.song import  SongUpdateSchema
from schemas.errors import *
import services.songs as songs_service
import services.ratings as ratings_service


songs = Blueprint(name="songs", import_name=__name__)

# post song
@songs.route('/', methods=['POST'])
@login_required
def post_song():
    """
    ---
    post:
      description: Creating a song
      requestBody:
        required: true
        content:
            application/json:
                schema: SongCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    print(request.json, "request.json")
    return songs_service.create_song(request.json)

# get song by id
@songs.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.get_song(id)

# update song 
@songs.route('/<id>', methods=['PUT'])
@login_required
def put_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    print(request.json, "request.json")
    return songs_service.update_song(id, request.json)

# delete song
@songs.route('/<id>', methods=['DELETE'])
@login_required
def delete_song(id):
    """
    ---
    delete:
      description: Deleting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.delete_song(id)

# get all songs
@songs.route('/', methods=['GET'])
@login_required
def get_songs():
    """
    ---
    get:
      description: Getting all songs
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Songs
            application/yaml:
              schema: Songs
        '401':

          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
    """
    return songs_service.get_songs()

# get all ratings for a song
@songs.route('/<id>/ratings', methods=['GET'])
def get_ratings(id):
    """
    ---
    get:
      description: Getting all ratings for a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Ratings
            application/yaml:
              schema: Ratings
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
    """
    return songs_service.get_ratings(id)

# create a rating for a song
@songs.route('/<id>/ratings', methods=['POST'])
def post_rating(id):
    """
    ---
    post:
      description: Creating a rating for a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
          - ratings
    """
    return songs_service.create_rating(id, request.json)

# delete a rating for a song
@songs.route('/<id>/ratings/<rating_id>', methods=['DELETE'])
def delete_rating(id, rating_id):
    """
    ---
    delete:
      description: Deleting a rating for a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
          - ratings
    """
    return ratings_service.delete_rating(rating_id)

