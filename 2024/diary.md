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
* if-ok statements.
