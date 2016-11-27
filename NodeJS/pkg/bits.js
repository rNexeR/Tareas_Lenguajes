
exports.setBit = function(n, pos){
    n |= (1 << pos)
    return n
}

// Clears the bit at pos in n.
exports.clearBit = function(n, pos){
    var mask = ~(1 << pos)
    n &= mask
    return n
}

exports.hasBit = function(n, pos) {
    var val = n & (1 << pos)
    return (val > 0)
}