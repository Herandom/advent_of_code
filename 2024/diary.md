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


## 2024-12-06
FIRST TRY WITH REAL DATA AND REGEX ON TASK 1 BABY!
Let's not talk about task 2.

### Completed
* Day 3: 1 star

### Learnings
* Go regex syntax is using google re2, a google implementation of regex or something.

## 2024-12-07
Turns out the method for my final few attemps at D3.2 yesterday was sound, I just had messed up the final regex I wrote, so I trimmed a few inputs to many.

### Completed
* Day3: 1 star

### Learnings

