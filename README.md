jigsaw 
======
Puzzle consists of sixteen square plastic pieces with different shapes on each side
of the square. Shapes are arrows pointing in, arrows pointing out, a ball and a 
cross. Shapes are either on the outside or cut out (negative). 

      
**(See png-file for scan of all pieces)**


In this ascii representation, the shapes in 
parentheses () are the shapes that are cut out
instead of sticking out.

	   v      v      v      o  
	(<)  o (>)  < (o)  o (o)  > 
	  (^)    (+)    (+)    (^)

	   ^      ^      +      +  
	(>)  < (o)  > (+)  o (>)  < 
	  (o)    (v)    (o)    (v)  

	   ^      +      v      + 
	(>)  + (<)  o (o)  < (>)  o 
	  (o)    (o)    (+)    (^) 

	   v      o      o      o  
	(+)  o (o)  o (o)  o (o)  > 
	  (v)    (^)    (^)    (+)  


The program just solves the puzzle by brute forcing all possibilities 
until it finds a solution. This is one of the solutions:


	  (o)     o     (v)    (v)   
	(>)  o (o)  > (>)  + (+)  < 
	   v     (+)     o      ^  

	  (v)     +     (o)    (^)    
	(o)  > (>)  < (<)  > (>)  <  
	   +     (v)     v      o   

	  (+)     v     (v)    (o) 
	(o)  + (+)  o (o)  > (>)  o 
	   o     (v)     ^      o    

	  (o)     v     (^)    (o)  
	(>)  o (o)  o (o)  + (+)  <  
	   o     (+)     o      ^   


