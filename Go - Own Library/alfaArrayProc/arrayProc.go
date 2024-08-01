package alfaArrayProc

import (
	"fmt"
	"reflect"
	"sort"
)

type IterableFunc func(element any, index int) bool
type MapFunc func(element any, index int) any
type SortFunc func(any, any) bool
type ReduceFunc func(accumulator, currentElement any, index int) any

func init() {

}

func Filter(iterable any, filterFunc IterableFunc) any {
	fmt.Println("Filter running")

	iterableValue := reflect.ValueOf(iterable)

	switch iterableValue.Kind() {
	case reflect.Slice, reflect.Array:
		resultSlice := reflect.MakeSlice(iterableValue.Type(), 0, 0)

		for i := 0; i < iterableValue.Len(); i++ {
			elem := iterableValue.Index(i).Interface()
			if filterFunc(elem, i) {
				resultSlice = reflect.Append(resultSlice, reflect.ValueOf(elem))
			}
		}

		return resultSlice.Interface()

	case reflect.Map:
		resultMap := reflect.MakeMap(iterableValue.Type())

		keys := iterableValue.MapKeys()
		for index, key := range keys {
			elem := iterableValue.MapIndex(key).Interface()
			if filterFunc(elem, index) {
				resultMap.SetMapIndex(key, reflect.ValueOf(elem))
			}
		}

		return resultMap.Interface()

	default:
		panic("iterable must be a slice, array, or map")
	}
}

func Sort(iterable any, lessFunc SortFunc) {
	iterableValue := reflect.ValueOf(iterable)

	switch iterableValue.Kind() {
	case reflect.Slice, reflect.Array:
		if iterableValue.Kind() == reflect.Slice {
			sort.Slice(iterable, func(i, j int) bool {
				elemI := iterableValue.Index(i).Interface()
				elemJ := iterableValue.Index(j).Interface()
				return lessFunc(elemI, elemJ)
			})
		} else {
			// If it's an array, convert it to a slice and then sort
			slice := reflect.MakeSlice(reflect.SliceOf(iterableValue.Type().Elem()), iterableValue.Len(), iterableValue.Len())
			reflect.Copy(slice, iterableValue)
			sort.Slice(slice.Interface(), func(i, j int) bool {
				elemI := slice.Index(i).Interface()
				elemJ := slice.Index(j).Interface()
				return lessFunc(elemI, elemJ)
			})
			reflect.Copy(iterableValue, slice)
		}

	case reflect.Map:
		// If it's a map, extract keys to a slice, sort the slice, and then create a new map with sorted keys and values
		keys := iterableValue.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			keyI := keys[i].Interface()
			keyJ := keys[j].Interface()
			return lessFunc(keyI, keyJ)
		})

		// Create a new map with the sorted keys and values
		sortedMap := reflect.MakeMapWithSize(iterableValue.Type(), len(keys))
		for _, key := range keys {
			value := iterableValue.MapIndex(key)
			sortedMap.SetMapIndex(key, value)
		}

		// Copy values from the sorted map to the original map
		for _, key := range keys {
			iterableValue.SetMapIndex(key, sortedMap.MapIndex(key))
		}

	default:
		panic("iterable must be a slice, array, or map")
	}
}

func Map(iterable any, mapFunc MapFunc) any {
	iterableValue := reflect.ValueOf(iterable)

	if iterableValue.Kind() != reflect.Slice && iterableValue.Kind() != reflect.Array {
		panic("iterable must be a slice or an array")
	}

	resultSliceType := reflect.SliceOf(reflect.TypeOf(mapFunc(iterableValue.Index(0).Interface(), 0)))
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

	for i := 0; i < iterableValue.Len(); i++ {
		elem := iterableValue.Index(i).Interface()
		result := mapFunc(elem, i)
		resultSlice = reflect.Append(resultSlice, reflect.ValueOf(result))
	}

	return resultSlice.Interface()
}

func Reduce(iterable any, reduceFunc ReduceFunc, initialValue any) any {
	iterableValue := reflect.ValueOf(iterable)

	if iterableValue.Kind() != reflect.Slice && iterableValue.Kind() != reflect.Array {
		panic("iterable must be a slice or an array")
	}

	accumulator := reflect.ValueOf(initialValue)

	for i := 0; i < iterableValue.Len(); i++ {
		currentElement := iterableValue.Index(i).Interface()
		accumulator = reflect.ValueOf(reduceFunc(accumulator.Interface(), currentElement, i))
	}

	return accumulator.Interface()
}

func Some(iterable any, someFunc IterableFunc) bool {
	fmt.Println("Filter running")

	iterableValue := reflect.ValueOf(iterable)
	result := false

	switch iterableValue.Kind() {
	case reflect.Slice, reflect.Array:

		for i := 0; i < iterableValue.Len(); i++ {
			elem := iterableValue.Index(i).Interface()
			if someFunc(elem, i) {
				result = true
				break
			}
		}

		return result

	case reflect.Map:

		keys := iterableValue.MapKeys()
		for i, key := range keys {
			elem := iterableValue.MapIndex(key).Interface()
			if someFunc(elem, i) {
				result = true
				break
			}
		}

		return result

	default:
		panic("iterable must be a slice, array, or map")
	}
}
