package main
import ("fmt")

func main() {
  var first,second string
  fmt.Print("Type a word: ")
  fmt.Scanln(&first)
  var length1 = len([]rune(first))
  var length2 = len([]rune(second))
  var c int
  c=0
  for i := 0; i <length1; i++ {
	if(first[i]=='i'||first[i]=='I')
	{
		c++;
		for j:=i+1;j<length1;j++ {
			if(first[i]=='a'||first[i]=='A')
			{
				c++;
				for k:=j+1;k<length1;k++ {
					if(first[i]=='a'||first[i]=='A')
					c+=1;
					if(c==3)
					break;
				}
					
			}
			if(c==3)
					break;
		}

	}
	if(c==3)
		break;
  }

  if(c==3)
	fmt.Print("Found!");
	else 
	fmt.Print("Not Found!");

}