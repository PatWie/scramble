# Go implementation of Pseudo encrypt


An implementation that scrambles id and encodes them as a string, similar to
(Instagram/YouTube id encodings).
All functions are *bijective* functions (on the positive domain), ie. they will
shuffle all integers from int32 (resp. int64) with zero collisions and can be undone.


Implements the postgresql [pseudo encrypt](https://wiki.postgresql.org/wiki/Pseudo_encrypt).

```
Scramble32(x)

 x  | pseudo_encrypt
-----+----------------
 -10 |    -1270576520
  -9 |     -236348969
  -8 |    -1184061109
  -7 |      -25446276
  -6 |    -1507538963
  -5 |     -518858927
  -4 |    -1458116927
  -3 |     -532482573
  -2 |     -157973154
  -1 |    -1105881908
   0 |     1777613459
   1 |      561465857
   2 |      436885871
   3 |      576481439
   4 |      483424269
   5 |     1905133426
   6 |      971249312
   7 |     1926833684
   8 |      735327624
   9 |     1731020007
  10 |      792482838
 ```


For convenience, these numbers can be encoded as strings (Instagram/YouTube like id encodings).

```go
encoding := scramble.AlphabethEncoder32{Alphabeth: scramble.AlphaNumeric}
slug:= encoding.Encode( scramble.Scramble32(4) ) // returns "NTySG"
originalID:= Unscramble32(encoding.Decode(slug)) // gives 4
```

See `readme_test.go`
