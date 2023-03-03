# names

`*names.Names` implements a set of `string`s which can be searched by substrings.

The set is case-sensitive. To get case-insensitive behavior, apply `strings.ToLower()` or `strings.ToUpper()` on any strings passed and translate them back to their original form using a lookup `map[string]string` if needed.

## Example Usage

```golang
const maxLookupSubstringRuneCount = 3
const supportRemove = true
people := names.New(maxLookupSubstringRuneCount, supportRemove)

people.Add("amelia")
people.Add("james")
people.Add("maxim")
people.Add("luke")
people.Add("sam")
result := people.Find("am", nil) // []string{"amelia", "james", "sam"}
people.Contains("james") // true

people.Remove("james") // this panics IFF supportRemove is false
result = people.Find("am", result) // []string{"amelia", "sam"}; result reused and returned as result[:2]
people.Contains("james") // false
```

## Sorting

Often after querying for matches, one may want to sort the result alphabetically. To do this, see functions `SortCIFunc`, `NewSortUXFunc` and `NewSortUXCIFunc` for use with `slices.SortFunc()` on string slices returned from `Find()`.

## Technical details

This implementation provides fast lookup speeds by maintainig a `map[string][]string` which maps all unique substrings up to a specified maximum length of every added string to a list of potential matches. You specify said maximum substring length as the first parameter in a call to `names.New()`. Higher values mean faster lookup speed at the cost of higher memory usage and lower insertion speed. As such, this implementation performs great for short strings, but poorly for long strings (i.e. over 100 runes per string).

It is possible to remove strings from the set, but only if you pass `true` as the second parameter in a call to `names.New()`. This is because an extra lookup map needs to maintained to allow for fast removal, which however increases memory usage and insertion speed further.
