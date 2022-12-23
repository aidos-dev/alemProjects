# ascii-art-output 

#### ascii-art-output is a programm which transforms a string from the command line into ascii-art format and stores the result to a file

It recieves in the command line 3 arguments: 

1) any input that user would like to be converted into ascii-art fomat

2) a banner name that actually means a name of a font for the ascii-art format

3) a file name where the ascii-art output will be stored. A user may specify any name as .txt file

By specifying these 3 arguments in the following format: 

Usage: 
> go run . [STRING] [BANNER] [OPTION]

Example: 
> go run . something standard --output=<fileName.txt>

#### a user can transform to ascii-art form any string that he wants, also there are 3 fonts(banners) that a user can choose:

- standard
- shadow
- thinkertoy

After completing the command described above in order to display the result just write the next command:

> cat <fileName.txt>

![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)