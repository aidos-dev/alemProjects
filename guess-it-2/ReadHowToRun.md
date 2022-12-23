# guess-it-2

## guess-it-2 is a program that makes prediction for every next number in a data set in a form of range of numbers (from a lower limit to an upper limit)

### How to test the program:

1. Download zip file from here: https://assets.01-edu.org/guess-the-number.zip . 

2. Extract the content of the zip file

3. Clone this repository to your machine

4. Place all extracted files (from step 2 of this instruction) to the root directory of the cloned repository.

5. In the root directory run the command:

```
make run
```

6. After running step 5 you will see in a terminal window:

```
Listening in port 3000
```

7. In your browser go to:
```
http://localhost:3000/
```

8. After opening your browser of preference in the port [3000] http://localhost:3000/, if you try clicking on any of the `Test Data` buttons, you will notice that in the Dev Tool/ Console there is a message which tells you that you need another guesser besides the student.

Adding a guesser is simple. You just need to add in the URL a guesser, in other words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose which of the random data set to test. After that you can wait for the program to test all of the values (boooooring), or you can click `Quick` in order to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays clicking on the `Clean` button after each test.

#### If you terminated the local host and want to start it to test it agian, in a root derectory run the command:

```
node server.js
```

... after that you can follow the same instructions as in step 8

Thank you for testing my program!

![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)