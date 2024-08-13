package main

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"testing"
)


func TestPathTransformFunc(t *testing.T){
	key:="nomsbestpicture"
	pathKey:=CasPathTransform(key)
	fmt.Println(pathKey.Pathname)
	expectedOriginalKey:="29be35d95381eb59abda8a6c4e5edb456569736d"
	expectedPathName:="29be3/5d953/81eb5/9abda/8a6c4/e5edb/45656/9736d"
	if pathKey.Pathname!=expectedPathName{
		t.Errorf("have %s want %s", pathKey.Pathname, expectedPathName)
	}

if pathKey.Filename!=expectedOriginalKey{
		t.Errorf("have %s want %s", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T){
	opts:=StoreOpts{
		PathTransformFunc: CasPathTransform,
	}
	    s:=NewStore(opts)
		key:="nomsspecials"
		data:=[]byte("some jpg bytes")
		if err:=s.WriteStream(key,bytes.NewReader(data)); err!=nil{
			t.Error(err)
		}

		if err:=s.Delete(key);err!=nil {
			t.Error(err)
		}
	
}

func TestStore(t *testing.T){
	/*opts :=StoreOpts{
		PathTransformFunc:CasPathTransform /*DefaultPathTransformFunc*//*,
	}
	//s:=NewStore(opts)*/
	s:=newStore()

	defer teardown(t,s)

	for i:=0; i<50; i++ {

	key:=fmt.Sprintf("foo_%d", i)

	dataNoConvert:=[]byte("some jpg bytes")

	data:=bytes.NewReader(dataNoConvert)

	if err:=s.WriteStream(key, data); err!=nil{
		t.Error(err)
	}

    if ok:=s.Has(key);!ok {
		t.Errorf("expected to have key %s\n", key)
	}


	r,err:=s.Read(key)
	if err!=nil {
		t.Error(err)
	}

	b,_:=ioutil.ReadAll(r)

	if string(b)!=string(dataNoConvert){
		t.Errorf("want %s have %s ", string(dataNoConvert),string(b))
	}
	if err:=s.Delete(key); err!=nil {
		t.Error(err)
	}
	if ok:=s.Has(key);ok {//test if after delete folder exist
		t.Errorf("expected to not have key %s\n", key)
	}
}
	
}


func newStore() *Store {
	opts:=StoreOpts{
		PathTransformFunc: CasPathTransform,
	}
		return NewStore(opts)
}

func teardown(t *testing.T, s *Store){
	if err:=s.Clear(); err!=nil {
		t.Error(err)
	}
}
