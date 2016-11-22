package bits

// Sets the bit at pos in the integer n.
func SetBit(n byte, pos uint) byte{
    n |= (1 << pos)
    return n
}

// Clears the bit at pos in n.
func ClearBit(n byte, pos uint) byte{
    mask := ^(1 << pos)
    n &= byte(mask)
    return n
}

func HasBit(n byte, pos uint) bool {
    val := n & (1 << pos)
    return (val > 0)
}