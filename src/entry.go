package jpsplus

import (
	"fmt"
)

func PreprocessMap(bits []byte, w int, h int, filename string) {
	fmt.Printf("Writing to file '%s'\n", filename)
	precomputeMap := newPrecomputeMap(w, h, bits)
	precomputeMap.CalculateMap()
	precomputeMap.SaveMap(filename)
}

func PrepareForSearch(bits []byte,  w int, h int, filename string){
    fmt.Printf("Reading from file '%s'\n", filename)
    precomputeMap := newPrecomputeMap(w, h, bits)
    precomputeMap.LoadMap(filename)
    var preprocessedMap = []*JumpDistancesAndGoalBounds{}
    preprocessedMap = precomputeMap.GetPreprocessedMap()
    return (void*)new JPSPlus(preprocessedMap, bits, w, h)
}

bool GetPath(void *data, xyLoc s, xyLoc g, std::vector<xyLoc> &path)
{
    JPSPlus* search = (JPSPlus*)data;
    return search->GetPath((xyLocJPS&)s, (xyLocJPS&)g, (std::vector<xyLocJPS>&)path);
}