package log

type Option func(opts *options)

type options struct {
	filename  string
	maxSize   int
	maxBackup int
	maxAge    int
	compress  bool
	logType   string
	level     string
}

func WithFilename(filename string) Option {
	return func(opts *options) {
		opts.filename = filename
	}
}

func WithMaxSize(maxSize int) Option {
	return func(opts *options) {
		opts.maxSize = maxSize
	}
}

func WithMaxBackup(maxBackup int) Option {
	return func(opts *options) {
		opts.maxBackup = maxBackup
	}
}

func WithMaxAge(maxAge int) Option {
	return func(opts *options) {
		opts.maxAge = maxAge
	}
}

func WithCompress(compress bool) Option {
	return func(opts *options) {
		opts.compress = compress
	}
}

func WithLogType(logType string) Option {
	return func(opts *options) {
		opts.logType = logType
	}
}

func WithLevel(level string) Option {
	return func(opts *options) {
		opts.level = level
	}
}
