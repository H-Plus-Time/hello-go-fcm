# hello-go-fcm
All-in-one hello world example of Firebase Cloud Messaging via Go + Javascript.

Backend periodically publishes fake data to a series of user-specified topics. Additionally, it provides a simple REST API for a client to discover and request a subscription to a topic; This is necessary because subscribing to a topic requires a server key (api tokens ability to do this has been disabled).

