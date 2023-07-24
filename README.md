# SCOIR Code Challenge for Full Stack Engineers

This repo contains an exercise intended for Full Stack Software Engineers.

## Instructions

1. Click "Use this template" to create a copy of this repository in your personal github account.
1. Using technologies of your choice, complete [the assignment](./Assignment.md).
   - Implement a Front End UI for selecting dogs and a Back End API for saving a user's dog selections.
   - Our preference would be to see the Front End in [React](https://reactjs.org/) with the UI implemented using [MUI](https://mui.com/), and the Back End to be in [Go](https://go.dev).
1. Update this README with
   - a `How-To` section containing any instructions needed to run/access your system.
   - an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (code_challenge@scoir.com) with a link to your newly created repo containing the completed exercise.

## Expectations

1. This exercise is meant to drive a conversation between you and Scoir's hiring team.
1. Please invest only enough time needed to demonstrate your approach to problem solving and code design.
1. Within reason, treat your solution as if it would become a production system.
1. If you have any questions, feel free to contact us at code_challenge@scoir.com

## How-To

How to start up the server and the react frontend

### Server

cd into scoir-server directory
docker build -t your-image-name .
docker run -p 8080:8080 your-image-name

## Assumptions

TODO reword this
I have designed the system layout expectign this to be a small-mid sized project. Some things I have done because of this is organizing the server by top-level feature as opposed to functionality. This should be more ideal for a small project such as this. I also grouped util and config in main space since it is such a small project.
