fpo: a dummy image server.

Requests are parsed: http://mydomain:4000/[height]x[width]/[color]
[color] can be either: hex value for the desired color, 'random', or omitted for a standard gray.

examples: 
	http://mydomain:4000/200x300/ffffaa
	http://mydomain:4000/100x400/random
	http://mydomain:4000/250x3000/
