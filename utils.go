package main

import "strings"

func matchesWildcard(origin, pattern string) bool {

    if !strings.HasPrefix(pattern, "*.") {
        return false
    }

    base := strings.TrimPrefix(pattern, "*.") 

    return strings.HasSuffix(origin, base)
}

var _ = matchesWildcard