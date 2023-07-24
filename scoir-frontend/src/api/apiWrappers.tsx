import {getDogs, getRandomDog, getDogByBreed, addDog, removeDog, clearDogs} from './api'
import noop from 'noop-ts';

export async function loadDogs(token: string){
    const dashboard = await getDogs(token);

      // update local storage
      if ('message' in dashboard) {
        localStorage.setItem('dogs', JSON.stringify(dashboard['message']));
      }
}

export async function addNewDog(token: string, breed?: string){
    var newDog: any
    if (breed) {
        newDog = await getDogByBreed(
            token,
            breed
          );
    } else {
        newDog = await getRandomDog(
            token
        );
    }

    if (!('message' in newDog)) {
        // TODO should error
        return 
    }
    await addDog(token, newDog['message'])
    await loadDogs(token)
}

export async function deleteDog(token: string, picture: string) {
    await removeDog(token, picture)
    await loadDogs(token)
}

export async function clearAllDogs(token: string){
    await clearDogs(token)
    await loadDogs(token)
}