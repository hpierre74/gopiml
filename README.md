Go application using Machine learning through **facebox**, a [machinebox.io](machinebox.io) facial recognition tool and [gocv](#), a library offering bindings for [OpenCV](#), which leverages image parsing in videos.

## Features

- A web client written in [react](#).

  - To upload image files to the server
  - To configure voice responses _coming soon_
  - To configure actions to be taken _coming later but soon_

- An Image API written in GO

  - Serve the web client
  - Process Post requests
  - Detect faces in picture
  - Crop faces from pictures
  - Teach them to Facebox

- A Raspberry Pi Webcam Application written in GO
  - To recognize faces taught through the web client
  - To execute actions triggered by voice (send sms, mails, add to shopping-list, etc...)

## Roadmap

### Web

Requirements

- :ok_hand: Host Facebox on Google Cloud.
- :ok_hand: Create a web client form to upload images
- :ok_hand: Create a web server to receive images from the client
- :point_right: Check and Teach images to Facebox
- :point_right: Convert and store images in Base64 on Google Storage
- :point_right: Add a programmable action interface
- :point_right: Add a general config interface
- :point_right: Find web hosting
  - :point_right: Firebase : would mean dropping Go SPA web server solution
  - :point_right: Google AppEngine : Costs ?
- :point_right: Find cheaper facebox hosting

### Machine

- :ok_hand: Create a program to recognize taught faces through webcam
- :ok_hand: Use google speech api to greet user when a face is matched
- :ok_hand: Program answers depending on user response to greeting
- :point_right: Program actions depending on user response
- :point_right: Use remote config to activate actions and industrialize process

## Contributing

If you too are interested in learning golang by doing funny/black-mirror-like things, feel free to fork and push your PR !
