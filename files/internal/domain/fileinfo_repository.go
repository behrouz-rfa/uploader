package domain

type FileInfoRepository interface {
	Upload(info FileInfo) error
}
