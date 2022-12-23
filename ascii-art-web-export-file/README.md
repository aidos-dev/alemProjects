# Ascii-art-web-export-file

### Ascii-art-web-export-file is a web-site which transforms a text to ascii-art format and displays the result. It is also possible to save your ascii-art to a text file

#### Additionaly the project has a Dockerfile for building a Docker image of a web-site. With help of this docker image a web site can run from the docker container.

### How to run the programm

1. Clone the repository to your machine
2. In the root directory of the cloned repository write the command:

```
make run
```

3. After completing the command above you will see the notice in the terminal window:

> " \<current date and time\> Listening on http://localhost:8080/ "

4. Hold the "Ctrl" key and Click on the link:

```
http://localhost:8080/
```

5. When you finish testing the web site you can terminate a local host with the command in a terminal window:

```
"Ctrl" + C
```

#### On the web-site a user can transform to ascii-art form any text that he wants with one of 3 fonts(banners):

- standard
- shadow
- thinkertoy

All you need to do is just input a desired text to the text input window on the web-site and click "Convert" button

After completing the command described above a given text will be displayed in ascii-art format in a chosen font

---

### How to run the web site from a docker container

1. In the root directory of the cloned repository write the command:

```
make docker
```

2. When you make sure that virtual machine is alive just run the command:

```
exit
```

3. To make sure that the web-site actually works, you can open your browser and go to:

```
http://localhost:8080/
```

---

### If you want to see the link to the web-site in a terminal window and test the web-site follow the next steps

1. Stop running container from the previous spteps with the command:

```
docker stop ascii-cont
```

2. In the root directory of the cloned repository write the command:

```
docker run -p 8080:8080 ascii-docker
```

3. After compliting the command you will see the notice in the terminal window: " \<current date and time\> Listening on http://localhost:8080/ "

4. Hold the "Ctrl" key and Click on the link:

```
http://localhost:8080/
```

---

### How to remove a docker container and a docker image from your machine after testing the web site

1. In the root directory of the cloned repository write the command:

```
make remove
```

2. After completing the command above you will see in your terminal window that a docker container and a docker image were removed.

---

### Thank you for testing my program!

---

##### Some basic docker commands

Check the list of existing docker images:

```
docker images
```

Check the list of existing docker containers:

```
docker ps -a
```

Stop running container:

```
docker stop ascii-cont
```

![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)
