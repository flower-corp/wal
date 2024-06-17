package wal

import "os"

type Option func(*Options)

// Options represents the configuration options for a Write-Ahead Log (WAL).
type Options struct {
	// DirPath specifies the directory path where the WAL segment files will be stored.
	DirPath string

	// SegmentSize specifies the maximum size of each segment file in bytes.
	SegmentSize int64

	// SegmentFileExt specifies the file extension of the segment files.
	// The file extension must start with a dot ".", default value is ".SEG".
	// It is used to identify the different types of files in the directory.
	// Now it is used by rosedb to identify the segment files and hint files.
	// Not a common usage for most users.
	SegmentFileExt string

	// Sync is whether to synchronize writes through os buffer cache and down onto the actual disk.
	// Setting sync is required for durability of a single write operation, but also results in slower writes.
	//
	// If false, and the machine crashes, then some recent writes may be lost.
	// Note that if it is just the process that crashes (machine does not) then no writes will be lost.
	//
	// In other words, Sync being false has the same semantics as a write
	// system call. Sync being true means write followed by fsync.
	Sync bool

	// BytesPerSync specifies the number of bytes to write before calling fsync.
	BytesPerSync uint32
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var DefaultOptions = Options{
	DirPath:        os.TempDir(),
	SegmentSize:    GB,
	SegmentFileExt: ".SEG",
	Sync:           false,
	BytesPerSync:   0,
}

// WithDirPath sets the directory path where the WAL segment files will be stored.
func WithDirPath(dir string) Option {
	return func(o *Options) {
		o.DirPath = dir
	}
}

// WithSegmentSize sets the maximum size of each segment file in bytes.
func WithSegmentSize(size int64) Option {
	return func(o *Options) {
		o.SegmentSize = size
	}
}

// WithSegmentFileExt sets the file extension of the segment files.
func WithSegmentFileExt(ext string) Option {
	return func(o *Options) {
		o.SegmentFileExt = ext
	}
}

// WithSync sets the whether to synchronize writes through os buffer cache and down onto the actual disk.
func WithSync(sync bool) Option {
	return func(o *Options) {
		o.Sync = sync
	}
}

// WithBytesPerSync sets the number of bytes to write before calling fsync.
func WithBytesPerSync(bytesPerSync uint32) Option {
	return func(o *Options) {
		o.BytesPerSync = bytesPerSync
	}
}
