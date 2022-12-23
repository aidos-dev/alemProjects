# ascii-art-web-dockerize

### ascii-art-web-dockerize is a project that builds a Docker image of a web-site with help of a Dockerfile

### How to run the programm:

1. Clone the repository to your machine
2. In the root directory of the cloned repository write the command: 

> make run

3. When you make sure that virtual machine is alive just run the command: 

> exit

4. To make sure that the web-site actually works, you can open your browser and go to:

> http://localhost:8080/

### If you want to see the link to the web-site in a terminal window and test the web-site follow the next steps:

1. Stop running container from the previous spteps with the command:

> docker stop ascii-cont

2. In the root directory of the cloned repository write the command:

> docker run -p 8080:8080 ascii-docker

3. After compliting the command you will see the notice in the terminal window: " \<current date and time\> Listening on http://localhost:8080/ "

4. Hold the "Ctrl" key and Click on the link: 

> http://localhost:8080/


#### On the web-site a user can transform to ascii-art form any text that he wants with one of 3 fonts(banners):

- standard
- shadow
- thinkertoy

All you need to do is just input a desired text to a text input window on the web-site and click "Convert" button 


After completing the command described above a given text will be displayed in ascii-art format in a chosen font


![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)