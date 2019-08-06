Go application using Machine learning through **facebox**, a [machinebox.io](machinebox.io) facial recognition tool and [gocv](#), a library offering bindings for [OpenCV](#), which leverages image parsing in videos.

## Features

- A web client written in [react](#).
  - To upload image files to the server
  - To configure voice responses _coming soon_
  - To configure actions to be taken _coming later but soon_

# Roadmap

### - Web

Requirements

[ ] Host Facebox on Google Cloud.
[ ] Create a web client form to upload images
[ ] Create a web server to receive images from the client
[ ] Check and Teach images to Facebox
[ ] Convert and store images in Base64 on Google Storage
[ ] Add program action configuration interface (Firebase ?)

### - Machine

[ ] Create a program to recognize taught faces through webcam
[ ] Use google speech api to greet user when a face is matched
[ ] Program answers depending on user response to greeting
[ ] Program actions depending on user response
[ ] Use remote config to activate actions
