def SetBit(n, pos):
    n |= (1 << pos)
    return n

def ClearBit(n, pos):
    mask = ~(1 << pos)
    n &= (mask)
    return n

def HasBit(n, pos):
    val = n & (1 << pos)
    return (val > 0)