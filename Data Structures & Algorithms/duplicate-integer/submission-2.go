// if we use a hash map to store the frequency of each number ( or just a simple bool in this case )
func hasDuplicate(nums []int) bool {
    lookup := make(map[int]bool)
    for _ ,v := range nums {
        if lookup[v]{
            return true
        }
        lookup[v] = true
    }
    return false
}
