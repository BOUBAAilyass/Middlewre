<template>
  <div style="width: 100%">

    <br/>
    <br/>
    <h2 style="padding: 0 15px">My ratings</h2>
    <div v-if="userRatings.ratings.length === 0" style="margin: 0 auto">
      <h3>No rating to display :-(</h3>
    </div>
    <div class="row" style="margin: 0">
      <div v-for="rating in userRatings.ratings" v-bind:key="rating.id" class="col-sm-4" style="padding: 10px">
        <div class=" card" style="margin: 10px 0">
          <h3 class="card-header">{{ rating.music_id }} : {{ rating.rating }} / 5 <span class="fas fa-star" style="color: #ffff00"/></h3>
          <div class="card-body">
            <ul>
              <li><span style="color: #888">Content:</span> {{ rating.content }}</li>
              <li><span style="color: #888">Note:</span> {{ rating.rating }} / 5 <span class="fas fa-star" style="color: #ffff00"/></li>
              <li><span style="color: #888">Date:</span> {{ rating.date }}</li>
               <li><span style="color: #888">Technical details:</span></li>
              <ul>
               <li><span style="color: #888">Rating ID:</span> {{ rating.id }}</li>
              </ul>
            </ul>
          </div>
          <div class="card-footer" style="text-align: right">
            <span v-on:click="deleteComment(rating)" class="fas fa-times" style="cursor: pointer; color: #ff0000"/>
          </div>
        </div>
      </div>
    </div>

  </div>

</template>

<script setup>
import {onMounted, reactive} from "vue";
import {useAxios} from "@vueuse/integrations/useAxios";
import {useAuthStore} from "@/stores/auth_store.js";
import {storeToRefs} from "pinia";
import {useGeneralResponses} from "@/composables/general_responses.js";
import {useToast} from "vue-toast-notification";

import Vue3WaveAudioPlayer from 'vue3-wave-audio-player'


const userRatings = reactive({
  ratings: []
})
const songs = reactive({
  songs: []
})

const authStore = useAuthStore()
const {currentUser} = storeToRefs(authStore)

const generalResponses = useGeneralResponses()
const toast = useToast();

onMounted(() => {
  
  getRatings()
  
})

// Done

async function deleteComment(data) {
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'DELETE',
  }
  const {error} = await useAxios(authStore.authBaseUrl + '/ratings/' + data.id, config)
  if (!error.value) {
    toast.success("Rating deleted")
    await getRatings()
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}


async function getSongs() {
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'GET',
  }
  const {data, error} = await useAxios(authStore.authBaseUrl + 'songs/', config)

  if (!error.value) {
    
   songs.songs = []
   songs.songs = data.value
   
      
  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}
async function getRatings() {
  
  
  const config = {
    headers: authStore.authAxiosConfig,
    method: 'GET',
  }
 
  const {data, error} = await useAxios(authStore.authBaseUrl + 'users/'+authStore.currentUser.id+'/ratings', config)


  if (!error.value) {
   

  await getSongs()
   

    userRatings.ratings = []

    data.value.forEach(e => {
      console.log(e)
        songs.songs.forEach(s => {
          console.log(s.title)
          if(s.id === e.music_id){
            console.log(s.title)
            
            e.music_id = s.title
            userRatings.ratings.push(e)
          }

    })
    })

    

  } else {
    generalResponses.manageError(error.value)
    // manage error and let the component display it however it wants to
    return Promise.reject(error.value)
  }
}



</script>
