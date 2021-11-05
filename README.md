### About the project
<hr />
<p>
An application designed to send emails using a STMP server of your choice, without the need for extreme settings.
In order for the user to be able to choose which model to use, the project applies two modes of send the emails.

 The first way is from an API available within the application
where the user gets through two routes, send a template that will be fired

The application also has an implementation for messaging using Apache Kafka
 where he listens to a topic and sends the emails present in it. Message sended to topic must follow API Enconded Endpoint or it will not work. 
 </p>

 
### Run with Docker
<hr />
<p> 
 You can get more information about image on <a href="https://hub.docker.com/repository/docker/aawadallak/email-dispatcher" target="_blankpage">DockerHub</a>
</p>

```
    $ docker run --name email -p 5005:6728 -e KAFKA_BROKERS=kafkabrokers \
     -e KAFKA_TOPIC_READER=yourtopic \
     -e SMTP_SERVER=server \
     -e SMTP_USER=user@example.com \
     -e SMTP_PASS=pass \
     -e SMTP_PORT=port \
    -d aawadallak/email-dispatcher
```

## Endopoints API

<p><strong>Dispatch using base64</strong></p>
Attachments must contains enconded base64 content or will return an error. 

 - Cc is an optional parameter
 - Body is an optional parameter

```
Method: POST
Path: http://localhost:{exposed_port}/api/v1/email/dispatch
Body: content-type application/json
```

![enter image description here](https://user-images.githubusercontent.com/74802742/140552332-355dd506-ee3d-401f-81d4-8aaf24a8f88c.png)

<p><strong>Dispatch using multipart files</strong></p>

```
Method: POST
Path: http://localhost:{exposed_port}/api/v1/email/dispatch/attachment
Body: content-type multipart/form-data
```

 - Cc is an optional parameter
 - Body is an optional parameter


![enter image description here](https://user-images.githubusercontent.com/74802742/140552758-a25ea1fd-7bf1-4274-8d38-6cee4d139262.png)

<p><strong>Response</strong></p>
<p>Because process of sending email is async, response is only about if request is correct, not about if email was t correctly</p>

![enter image description here](https://user-images.githubusercontent.com/74802742/140553325-306e6dec-d9cb-4eb2-b32a-5c2207c0ff0c.png)

### 💪 How to contribute to the project
```
1. Do a **fork** of the project.
2. Create a new branch with your changes: `git checkout -b my-feature`
3. Save your changes and create a commit with a message: `git commit -m "feature: My new feature"`
4. Send your changes: `git push origin my-feature`
```

