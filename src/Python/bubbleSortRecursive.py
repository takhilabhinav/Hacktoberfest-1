def recursive_bubble_sort(my_list): 
	for i, num in enumerate(my_list): 
		try: 
			if my_list[i+1] < num: 
				my_list[i] = my_list[i+1] 
				my_list[i+1] = num 
				recursive_bubble_sort(my_list) 
		except IndexError: 
			pass
	return my_list 

my_list = [19,2,31,45,6,11,121,27]
print ("List before sort:", my_list)
recursive_bubble_sort(my_list) 

print("Sorted list:", my_list); 