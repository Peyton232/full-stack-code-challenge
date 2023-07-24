const baseURL = "http://127.0.0.1:8080"

export async function getRandomDog(token: string) {
    try {
      const response = await fetch(`${baseURL}/dogs/random`, { 
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}` 
        },
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }

  export async function getDogByBreed(token: string, breed: string) {
    try {
      const response = await fetch(`${baseURL}/dogs/${breed}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }
  

  export async function loginUser(username: string, password: string) {
    try {
      const response = await fetch(`${baseURL}/login`, { 
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }

  export async function getDogs(token: string) {
    try {
      const response = await fetch(`${baseURL}/users/all`, { 
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}` 
        },
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }

  export async function clearDogs(token: string) {
    try {
      const response = await fetch(`${baseURL}/users/clear`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }
  
  export async function addDog(token: string, dog: string) {
    try {
      const response = await fetch(`${baseURL}/users/add`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          picture: dog,
        }),
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }
  
  export async function removeDog(token: string, picture: string) {
    try {
      const response = await fetch(`${baseURL}/users/remove`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ picture }), // Send the picture to be removed in the request body
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty JSON object on error
      return {};
    }
  }
  
  export async function getAllBreeds(token: string) {
    try {
      const response = await fetch(`${baseURL}/dogs/breed/list/all`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        // If the response status is not OK (2xx), throw an error
        throw new Error('Network response was not ok');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error during fetch:', error);
      // Return an empty array on error
      return [];
    }
  }
  