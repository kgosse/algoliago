# algoliago
The test subject can be found: [here](https://gist.github.com/sfriquet/55b18848d6d58b8185bbada81c620c4a)

## Time Complexity

My focus was to organise the data so that I could have the best possible response time to queries.
My program can answer to "count" queries in constant time, O(1). For "popular" queries, the complexity
is O(n), n being the size parameter. Actually, I can also make it in O(1) if I don't loop through each
popular querie to construct the response. The idea would be to have them in the response format already and to directly
pick them. But I'm assuming the size parameter in not supposed to be that big, so this can make it.

## Data Structures

The challenge was to have as input a date and to respond the quickest way possible (in constant time). The date input can be a year or year+month or
year+month+day or year+month+day+hour+minute or even year+month+day+hour+minute:second

So if the client gives me a year and wants to know the distinct count,
I want to take that year input, look inside a Map data structure (the key
being the year) and return directly the "disctinct count". That's why
I've created a year structure with all the relevant informations. Then I have a Map of years.

Each year holds 12 months (no surprise here). So inside a year structure,
I have an array of 12 months. A month is another structure, like years. When given a year+month date, I can get the year, then the month and return directly the "distinct count".

Then, inside a month structure, we can find an array of 31 days. In a day structure, we have an array of 24 hours. The same logic continues till a I get to seconds.

This way of organizing the data helps me to find directly the answer in constant time.