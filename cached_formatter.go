package main


// CachedFormatter takes something that implements the Formatter interface, and adds
// a caching layer to it. It makes the assumption that that something's
// "Format(string) string" method is mathematically functional, in that the same
// input always returns the same output. (Do not use CachedFormatter if that is
// not true.) Based on that assumption, it caches the result from that something's
// "Format(string) string" method and uses that cache from its own "Format(string) string"
// method.
type CachedFormatter struct {
	formatter Formatter
	cache     map[string]string
}


// NewCachedFormatter creates a new CachedFormatter. The "formatter" parameter is Formatter
// to add a caching later to. The "cache_capacity_hint" parameter is a hint on how much
// capacity to initialize the cache with.
func NewCachedFormatter(formatter Formatter, cache_capacity_hint int) (*CachedFormatter) {
	cache := make(map[string]string, cache_capacity_hint)

	me := CachedFormatter{
		formatter: formatter,
		cache:     cache,
	}

	return &me
}


// CachedFormat does all the work for the "Format(string) string" method.
// From the outside, it is just like "Format(string) string" except it returns
// an additional bool return value that tells you if there was a cache hit or
// cache miss. A value of "true" means it was in the cache, and thus there was
// a cache hit. A value of "false" means it was not in the cache, and thus
// there was a cache miss.
func (me *CachedFormatter) CachedFormat(s string) (string, bool) {

	// Try to get it from the cache (if it is in there).
	// If it is not in the cache, then call "me.formatter.Format(string) string"
	// and then cache it.
		var was_in_cache bool
		var formatted    string

		if cached,ok := me.cache[s] ; ok {
			was_in_cache = true

			formatted = cached
		} else {
			was_in_cache = false

			formatted = me.formatter.Format(s)

			me.cache[s] = formatted
		}

	// Return.
		return formatted, was_in_cache
}


// Format implements the the "Format(string) string" method from the "Formatter" interface.
func (me *CachedFormatter) Format(s string) string {
	formatted, _ := me.CachedFormat(s)

	return formatted
}
