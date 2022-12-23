# ascii-art-web

### ascii-art-web is a programm which transforms a string from the web-site input form to ascii-art format and displays the result in a terminal window 


#### a user can transform to ascii-art form any string that he wants with one of 3 fonts(banners):

- standard
- shadow
- thinkertoy

The programm recieves in the command line 3 arguments: 

1) any input that user would like to be converted into ascii-art fomat

2) a banner name that actually means a name of a font for the ascii-art format

3) a flag with disired alignment


...in the following format: 

Usage: 
> go run . [STRING] [BANNER] [OPTION]

Example: 
> go run . something standard --align=right


After completing the command described above a given string will be displayed in the terminal window in ascii-art format in a chosen font and alignment


![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)