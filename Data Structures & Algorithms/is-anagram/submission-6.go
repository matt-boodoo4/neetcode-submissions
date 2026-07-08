// create a map of each character and how many times it occurs 
// populate the map with data 
// loop through both maps and look for each letter in the other map and check the count. 
// if the letter does not exist or the counts do not match, return false
func isAnagram(s string, t string) bool {
    // we can create a lookup to see if the characters appear the same amount of times in each string
    s1Map := make(map[rune]int, len(s))
    t1Map := make(map[rune]int, len(t))
    
    //step 1 index the runes of s
    for _ , v := range s {
        s1Map[v] += 1
    }
    for _ , v := range t {
        t1Map[v] +=1
    }
    for k, v := range s1Map {
        if count , ok := t1Map[k]; ok {
            if count != v {
                return false // mismatch 
            }
        } else {
            return false 
        }
    }
    
    for k, v := range t1Map {
        if count , ok := s1Map[k]; ok {
            if count != v {
                return false // mismatch 
            }
        } else {
            return false 
        }
    }
    return true
}
