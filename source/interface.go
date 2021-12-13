package source

import "io"

// ResourceClient defines the API interface to interact with source.
type ResourceClient interface {

	// GetContentLength get length of resource content
	// return source.UnknownSourceFileLen if response status is not StatusOK and StatusPartialContent
	GetContentLength(request *Request) (int64, error)

	// IsSupportRange checks if resource supports breakpoint continuation
	// return false if response status is not StatusPartialContent
	IsSupportRange(request *Request) (bool, error)

	// IsExpired checks if a resource received or stored is the same.
	// return false and non-nil err to prevent the source from exploding if
	// fails to get the result, it is considered that the source has not expired
	IsExpired(request *Request, info *ExpireInfo) (bool, error)

	// Download downloads from source
	Download(request *Request) (io.ReadCloser, error)

	// DownloadWithExpireInfo download from source with expireInfo
	DownloadWithExpireInfo(request *Request) (io.ReadCloser, *ExpireInfo, error)

	// GetLastModified gets last modified timestamp milliseconds of resource
	GetLastModified(request *Request) (int64, error)
}
