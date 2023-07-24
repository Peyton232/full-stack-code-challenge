import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import CameraIcon from '@mui/icons-material/PhotoCamera';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import CssBaseline from '@mui/material/CssBaseline';
import Grid from '@mui/material/Grid';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Link from '@mui/material/Link';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { useState } from 'react';
import {addNewDog} from '../apiWrappers'

interface Dog {
  Breed: string;
  Picture: string;
}

const options = ["affenpinscher", "african", "airedale", "akita", "appenzeller", "australian", "basenji", "beagle", "bluetick", "borzoi", "bouvier", "boxer", "brabancon", "briard", "buhund", "bulldog", "bullterrier", "cattledog", "chihuahua", "chow", "clumber", "cockapoo", "collie", "coonhound", "corgi", "cotondetulear", "dachshund", "dalmatian", "dane", "deerhound", "dhole", "dingo", "doberman", "elkhound", "entlebucher", "eskimo", "finnish", "frise", "germanshepherd", "greyhound", "groenendael", "havanese", "hound", "husky", "keeshond", "kelpie", "komondor", "kuvasz", "labradoodle", "labrador", "leonberg", "lhasa", "malamute", "malinois", "maltese", "mastiff", "mexicanhairless", "mix", "mountain", "newfoundland", "otterhound", "ovcharka", "papillon", "pekinese", "pembroke", "pinscher", "pitbull", "pointer", "pomeranian", "poodle", "pug", "puggle", "pyrenees", "redbone", "retriever", "ridgeback", "rottweiler", "saluki", "samoyed", "schipperke", "schnauzer", "segugio", "setter", "sharpei", "sheepdog", "shiba", "shihtzu", "spaniel", "spitz", "springer", "stbernard", "terrier", "tervuren", "vizsla", "waterdog", "weimaraner", "whippet", "wolfhound"];

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://mui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const defaultTheme = createTheme();

export default function Dashboard() {

  const [selectedBreed, setSelectedBreed] = useState<string | null>(null);

  let dogsData: Dog[] = [];

  const localStorageData = localStorage.getItem('dogs');
  if (localStorageData) {
    try {
      const parsedData = JSON.parse(localStorageData);
      if (Array.isArray(parsedData)) {
        dogsData = parsedData;
      }
    } catch (error) {
      console.error('Error parsing dogs data from localStorage:', error);
    }
  }

  const handleSearch = () => {
    var token = localStorage.getItem('token') || ""
    if (selectedBreed) {
      addNewDog(token, selectedBreed)
    }
  };

  const handleRandomDog = () => {
    var token = localStorage.getItem('token') || ""
    addNewDog(token)
  }

  return (
    <ThemeProvider theme={defaultTheme}>
      <CssBaseline />
      <AppBar position="relative">
        <Toolbar>
        </Toolbar>
      </AppBar>
      <main>
        <Box
          sx={{
            bgcolor: 'background.paper',
            pt: 8,
            pb: 6,
          }}
        >
          <Container maxWidth="sm">
            <Typography
              component="h1"
              variant="h2"
              align="center"
              color="text.primary"
              gutterBottom
            >
              Dog Catcher
            </Typography>
            <Stack
              sx={{ pt: 4 }}
              direction="row"
              spacing={2}
              justifyContent="center"
            >
              <Autocomplete
                disablePortal
                id="combo-box-demo"
                options={options}
                value={selectedBreed}
                onChange={(event, newValue) => setSelectedBreed(newValue)}
                sx={{ width: 300 }}
                renderInput={(params) => <TextField {...params} label="Breed" />}
              />
              <Button variant="contained" onClick={handleSearch}>
                Search
              </Button>
            </Stack>
            <Stack
              sx={{ pt: 4 }}
              direction="row"
              spacing={2}
              justifyContent="center"
            >
              <Button variant="outlined" onClick={handleRandomDog}>+ Catch A Random Breed </Button>
            </Stack>
          </Container>
        </Box>
        <Container sx={{ py: 8 }} maxWidth="md">
          {dogsData.length === 0 ? (
            <Typography
              variant="h5"
              align="center"
              color="text.primary"
              gutterBottom
            >
              There are currently no breeds caught.
              Search above to catch some!   
            </Typography>
          ) : (
            <Grid container spacing={4}>
              {dogsData.map((dog, index) => (
                <Grid item key={index} xs={12} sm={6} md={4}>
                  <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
                    <CardMedia
                      component="div"
                      sx={{
                        aspectRatio: '16/9',
                      }}
                    >
                      <img
                        src={dog.Picture}
                        alt={dog.Breed}
                        style={{
                          objectFit: 'cover',
                          width: '100%',
                          height: '100%',
                        }}
                      />
                    </CardMedia>
                    <CardContent sx={{ flexGrow: 1 }}>
                      <Typography gutterBottom variant="h5" component="h2">
                        {dog.Breed}
                      </Typography>
                    </CardContent>
                    <CardActions>
                      <Button size="small">Delete</Button>
                    </CardActions>
                  </Card>
                </Grid>
              ))}
            </Grid>
          )}
        </Container>
      </main>
      <Box sx={{ bgcolor: 'background.paper', p: 6 }} component="footer">
        <Typography variant="h6" align="center" gutterBottom>
          Footer
        </Typography>
        <Typography
          variant="subtitle1"
          align="center"
          color="text.secondary"
          component="p"
        >
        </Typography>
        <Copyright />
      </Box>
    </ThemeProvider>
  );
}