package main

import (
	//"bytes"
	//"crypto/md5"
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const defaultRootFolderName="ggnetwork"

func CasPathTransform(key string) Pathkey {

	hash:=sha1.Sum([]byte(key))
	hashStr:=hex.EncodeToString(hash[:])

	blockSize:=5
	slicelen:=len(hashStr)/blockSize

	paths:=make([]string, slicelen)

	for i:=0; i<slicelen; i++ {
         from, to:=i*blockSize, (i*blockSize)+blockSize
		 paths[i]=hashStr[from:to]
	}
	return Pathkey{
		Pathname:strings.Join(paths,"/") ,
		Filename: hashStr,
	}
	
}

type PathTransformFunc func(string) Pathkey

type Pathkey struct {
	Pathname string
	Filename string
}


func (p Pathkey)FirstPathName() string {
      paths:=strings.Split(p.Pathname, "/")
	if len(paths)==0 {
		return  "**"
	} 
	return paths[0]
}

func (p Pathkey) FullPath()string {

	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type StoreOpts struct {
	Root string
     PathTransformFunc  PathTransformFunc
}

type Store struct {
	// Root is the folder of the root, containing all the
	//folders/files of the system 
	
    StoreOpts StoreOpts
}

var DefaultPathTransformFunc = func (key string)Pathkey {
	return Pathkey{
		Pathname: key,
		Filename: key,
	}
}
func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc=DefaultPathTransformFunc
	}
	if len(opts.Root)==0 {
         opts.Root=defaultRootFolderName 
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(key string)bool {
	pathKey:=s.StoreOpts.PathTransformFunc(key)
    fullPathWithRoot:=fmt.Sprintf("%s/%s", s.StoreOpts.Root,pathKey.FullPath())
	fmt.Println(fullPathWithRoot)
	_,err:=os.Stat(fullPathWithRoot)
	return errors.Is(err, os.ErrNotExist)//verify if folder exist in the os système
}

func (s *Store) Clear() error {
	return os.RemoveAll(s.StoreOpts.Root)
}

func (s *Store) Delete(key string)error {
	pathKey:=s.StoreOpts.PathTransformFunc(key)
      
	defer func(){
		log.Printf("deleted [%s] from disk", pathKey.Pathname)
	} ()  
	
	/*if err:= os.RemoveAll(Pathkey.FullPath()); err!=nil {
		return err
	}*/
	firstPathNameWithRoot:=fmt.Sprintf("%s/%s", s.StoreOpts.Root,pathKey.FirstPathName())
	return os.RemoveAll(firstPathNameWithRoot)
}

func (s *Store) write(key string, r io.Reader) error {
	return s.WriteStream(key,r)
}

func (s *Store) Read(key string)(io.Reader, error){
	f,err:=s.ReadStream(key)
	if err!=nil {
		return nil, err
	}
	defer f.Close()
	buf:=new(bytes.Buffer)
	_,err=io.Copy(buf,f)

	return buf,err
}


func (s *Store) ReadStream(key string)(io.ReadCloser, error){
	pathKey:=s.StoreOpts.PathTransformFunc(key)
	fullPathWithRoot:=fmt.Sprintf("%s/%s", s.StoreOpts.Root,pathKey.FullPath())
	/*f, err:=os.Open(pathKey.FullPath())
	if err!=nil {
		return nil,err
	}
	return f,nil*/
	return os.Open(fullPathWithRoot)
} 


func (s *Store) WriteStream(key string, r io.Reader)error{

    pathKey:=s.StoreOpts.PathTransformFunc(key)
     pathNameWithRoot:=fmt.Sprintf("%s/%s", s.StoreOpts.Root,pathKey.Pathname)  
   if err :=os.MkdirAll( pathNameWithRoot, os.ModePerm); err!=nil {
	return err
   }

   buf:=new (bytes.Buffer)

   io.Copy(buf,r)

   //filenameBytes:=md5.Sum(buf.Bytes())
   //filename:=hex.EncodeToString(filenameBytes[:])

    //filename:="somefilename"
	//pathAndFilename:=pathKey.Pathname + "/" + filename
	fullPath:=pathKey.FullPath()
	fullPathWithRoot:=fmt.Sprintf("%s/%s", s.StoreOpts.Root,fullPath)
	f,err:=os.Create(fullPathWithRoot)

	if err!=nil {
		return err
	}
	//n,err:=io.Copy(f,buf)
	n,err:=io.Copy(f,r)
	if err!=nil {
		return err
	}
	log.Printf("written (%d) to disk with name (%s)", n, fullPathWithRoot)
	return nil

} 

/*func encode(gl any) ([]byte, error) {//boot dev fast encoding more than JSON
	
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(10)
	if err != nil {
		return nil,err
	}
	a:=network.Bytes()

	return a, nil
	
}*/


func decode(data []byte) (GameLog, error) {
	network:=bytes.NewReader(data)
	dec := gob.NewDecoder(network)
	var v GameLog
	err := dec.Decode(&v)
	if err != nil {
		return GameLog{},err
	}
	return v,nil
}

type GameLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}