# ascii-art-color

### ascii-art-color is a programm which transforms a string from the command line to ascii-art format and displays the result in a terminal window 


#### it is possible to color the output with the following colors: 

- gray
- red
- lime
- yellow
- dodger blue
- magenta
- cyan
- white
- black
- brown
- green
- orange
- blue
- purple
- light sea green
- silver


The programm recieves 3 arguments from the command line : 

1) any input that user would like to be converted into ascii-art fomat

2) a desired color from the list above

3) if necessary an index or a range of indexes of symbols wich a user wants to be colored


...in the following format: 

Usage: 
> go run . [STRING] [COLOR] [INDEX] 

Example: 
> go run cmd/main.go something --color=\<color\> --index=0:2 


After completing the command described above a given string will be displayed in the terminal window in ascii-art format in a chosen color

To test this programm you need to perform the following steps:

1. Clone this repository to your machine
2. In the root directory of the cloned repository writa a command specified above as an example


![black cat](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTlCmZNZTRr2c33iinneBtyyW2NjFkOSpGOLw&usqp=CAU)