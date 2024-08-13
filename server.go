package main

import "github.com/farisbrandone/distributed_file_storage/p2p"
type FileServerOpts struct{
	ListedAddr string
	StorageRoot string
	PathTransformFunc PathTransformFunc
	Transport  p2p.Transport
}
type FileServer struct {
	FileServerOpts
	store *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {

     storageOpts:=StoreOpts{
		Root:opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	 }

	return &FileServer{
		FileServerOpts:opts,
		store:NewStore(storageOpts),
	}
}

func (s *FileServerOpts) Start() error {
	if err:=s.Transport.ListenAndAccept(); err!=nil {
		return err
	}
	return nil
}

