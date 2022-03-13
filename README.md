# gob-decode-issue

I'm facing a weird issue in which the ReadCloser's `UnmarshalBinary` sets the write state, but the output of decode has extra stuff appended to it. 


To run

```
go run main.go
```