# Advent of Code 2024
## 2024-12-02
Did day1 and 2 in Python at work as some self learning, went quite smoothly (besides misreading the task)

At home, did day 1. File IO was the first hurdle but finally figured it out. The second hurdle was the fact that `sort.Ints` sorts values in place and doesn't have a return. That was painful and took way, way, way to long (like at least 30min+ if not more. Then it turns out that Go doesn't have an `abs(val int) int` function so I had to implement that one (not that it's difficult but still). Final hurdle was the fact that I went with maps (almost python dicts) and the syntax around that. if-ok solved the issue of inserting values from right by mistake, and then a second `value_count` map to make the multiplication factor correct...

### Completed
* Day 1: 2 stars
### Learnings
* File IO
* no int-abs func
* sort functions are in-place
* if-ok statements


## 2024-12-03
Started way to late and was distracted since I had stuff to deal with in that ran on a short timer. Stayed up way to late. Main challange was nothing really, just that vectorizing the Day 2 task is waaaay easier than iterating. I feel I can be smarter than this, but I did manage to skip one explicit for loop by calling Min and Max on the slice and multiplying, as well as using Contains.

I don't look forward to day 3 and the file reading. There will be a lot of string appending. Also regex. `No Wastl-dono! Yamete!` 
### Completed
* Day 2: 1 start
### Learnings
* The `slices` module from stdlib is cool! I used the `Contains`, `Min` and `Max` functions.
* `~` can not be used in an absolute path, they need to be truely absolute or relative.
* The `Min` * `Max` trick might actually be useful in other places!

## 2024-12-04
Late start again, but earlier than yesterday. Only targetting 2nd star of day 2 today, since day 3 is regex and day 4 is a word find. Both will require thinking. 
I successfully extracted the "is this thing ok" code from task 1, now it is just to brute force.

FIRST TRY WITH REAL DATA WAS SUCCESS BABY!
### Completed
* Day 2: 2nd star
### Learnings
* `slices.Delete` makes no sense. Don't use it, just handroll something instead. Addendum: or just make a copy of the slice and then Delete and remove n-indexes from the end. 
* There is an `...` operator! This should be understood!


## 2024-12-05
FIRST TRY WITH REAL DATA AND REGEX ON TASK 1 BABY!
Let's not talk about task 2.

### Completed
* Day 3: 1 star

### Learnings
* Go regex syntax is using google re2, a google implementation of regex or something.

## 2024-12-06
Turns out the method for my final few attemps at D3.2 yesterday was sound, I just had messed up the final regex I wrote, so I trimmed a few inputs to many.

For Day 4, a lot of issues with slices and slice-matrixes (matrixes from here on). Basically, I had to write funcions to go from `[]string` to `[][]string` and then write transpose, reverse row and "get diagnonals" functions. Finding the xmas/samx wasdone through regex.

Solution worked out on paper, to see the behaviour and keep track of all of the indexes, (indecies?)

Very late, but it's Friday and I took a 3-4 hour nap after work so it's probably OK. 

Second task went very well, did some more pen-and-paper work to get the index correct. Very happy that I thought to find the A and then just check the diags. Since I had the matrix generating code things went very smooth.

For both of Day 4 I did it on first try with real data!

### Completed
* Day3: 2nd star
* Day4: 2 stars

### Learnings
* I almost understand how to work with matrixes.

## 2024-12-07
Saturday! Day 5 task 1 was fun, but I think my solution might make task 2 more difficult. I also don't feel like writing a sorting algorithm so I will skip task 2 for now.

Day 6 basically worked as I had planned when reading the task at work, using a map to rotate and then just some boundry + `TheHinder` checks. Reminds me of the crystalization java lab, where I learned the hard way about switch/select case statements in Java not breaking, like in the surpringly sensible language VisualBasic

**OH NO! THE SWEDISH HINDER!** For Day 6.2 I had to rewrite the walk code a bit. I actually only kept track of the obstacles and the path inside the matrix, and tracked the guard walking direction outside of it. 
To see if we are in a loop, I tracked the rotations (ix, jx and post-rotation-direction). Each rotation after the initial 3 or 4 I checked if we had performed the same move once before. Doing a check for the path might have been possible as well, but I'm happy with my way of doing things :)

I did not bother figuring out how to duplicate `[][]string` and just read from file. I think the file reading probably slowed me down a lot. EDIT: tried, instead of reading, to just use the `[]string to [][]string` function for every hinder instead. There was no noticable performance gain. I think creating copies of the matrix could maybe speed things still, but to lazy to figure out how to implement. This task is however very parallizable, like a C-style OpenMP for the for loop would be so easy to get a speed boost.

There was also something strange in an if statement, I think I messed up a negation in an `false & false` statmenet. oops.
### Completed
* Day 5: 1st star
* Day 6: 2 stars
### Learnings
* File reading and shifting between two types of input in one file
* Slice intersections.Oh god python is easier in this regard.
* You don't need to return slices if you edit them inside of a function!

## 2024-12-08
I saw some AoC day 7 memes about recursion and `O(2^n)` and `O(4^ n)` scaling, so I got some help with figuring out day 7 part 1. 
This was needed, because I would not have figured out how to use recursion to generate all possible permutations of the operators without that. 
I am quite proud about my way of going about the "find all possible operation permutations with replacement". I used it if-ok way of doing thigns from day 1, and simply calculated them only once. 

One tricky part of the unput is that I probably missed one row since I used a dict to store all of the inputs and didn't account for 2 answers to the equation being the same. Oh well, I simply had to make my "check if this is OK" function a lot bigger.


### Commpleted 
* Day 7: 1 star

### Learnings
* I know to little about recursion and recursive algorithms.
* Maps to store some data is dangerous.




