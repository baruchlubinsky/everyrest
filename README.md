EveryREST
=========

Use this program as a server for Ember Data with the [REST Adapter](http://emberjs.com/guides/models/the-rest-adapter/). The server responds to all (WIP) requests from Ember Data as expected. All logic is handled by the Ember application.

Usage
==

Run locally using [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).

This project is ready to be deployed on [Google App Engine](https://cloud.google.com/appengine/docs/whatisgoogleappengine), and can be hosted for free - with limits more than sufficient for personal use.

An example of this webservice is running at [everyrest.appspot.com](http://everyrest.appspot.com). 

Run you own version by cloning this repository and changing the `application` attribute in each `.yaml` file. Rather than creating a multi-tenant application, let each user own their own server.


