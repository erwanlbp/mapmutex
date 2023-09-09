# Map of Mutexs

## The problem

I store a local `map[string]XXX` and have a function `GetXXXforResource(resource string) XXX` that fetches some data and stores it in the local `map`.

I want to make sure the fetching is done only once even if multiple concurrent calls occurs.

## The solution

Have a `map[string]sync.Mutex` and some methods `Lock(string)` / `Unlock(string)` on it.

## Install

```bash
go get github.com/erwanlbp/mapmutex
```
