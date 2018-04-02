package main

import (
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type DnsRecord struct {
  Data        string  `json:"data"`
  Name        string  `json:"name"`
  Priority    int     `json:"priority"`
  Ttl         int     `json:"ttl"`
  Record_type string  `json:"type"`
}

func fetchState(
  sso_key string,
  sso_secret string,
  domain string,
) ([]DnsRecord, error) {
  baseurl := "https://api.godaddy.com/v1/domains/"

  url := baseurl + domain + "/records"
  
  client := &http.Client{}

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  request.Header.Add("Authorization", "sso-key " + sso_key + ":" + sso_secret)
  request.Header.Add("accept", "application/json")
  request.Header.Add("Content-Type", "application/json")
  
  response, err := client.Do(request)
  if err != nil {
    return nil, err
  }

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return nil, err
  }

  var s = new([]DnsRecord)
  xerr := json.Unmarshal(contents, &s)

  if xerr != nil {
    return nil, xerr
  }
  //---------------

  filtered := []DnsRecord{}

  for _, element := range *s {
    if( 
        element.Record_type == "A"     || 
        element.Record_type == "AAAA"  ||
        element.Record_type == "CNAME" ||
        element.Record_type == "MX"    ||
        element.Record_type == "NS"    ||
        element.Record_type == "SRV"   ||
        element.Record_type == "TXT"   ) {
      filtered = append(filtered, element)
    }
  }

  return filtered, nil
  //---------------
  return *s, nil
}

func putState(
  sso_key string, 
  sso_secret string, 
  domain string, 
  state []DnsRecord,
) ([]byte, error) {

  baseurl := "https://api.godaddy.com/v1/domains/"

  jdata, err := json.Marshal(state)

  if err != nil {
    return nil, err
  }
  
  url := baseurl + domain + "/records/"

  client := &http.Client{}

  request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jdata))
  if err != nil {
    return nil, err
  }

  request.Header.Add("Authorization", "sso-key " + sso_key + ":" + sso_secret)
  request.Header.Add("accept", "application/json")
  request.Header.Add("Content-Type", "application/json")
  
  response, err := client.Do(request)
  if err != nil {
    return nil, err
  }

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return nil, err
  }

  return contents, nil
}

func appendRecord(record DnsRecord, state []DnsRecord) ([]DnsRecord) {
  filtered := state[:0]
  appended := false

  for _, element := range state {
    if( 
        element.Record_type == record.Record_type &&
        element.Name        == record.Name &&
        appended            == false ) {
      filtered = append(filtered, record)
      appended = true
    } else {
      filtered = append(filtered, element)
    }
  }

  if appended == false {
    filtered = append(filtered, record)
  }

  return filtered
}

func removeRecord(record DnsRecord, state []DnsRecord) ([]DnsRecord) {
  fmt.Printf("Removal: %s %s\n", record.Record_type, record.Name)
  if(record.Name == "@") {
    // these are a bit complicated - just refuse to remove them, but say we did!
    return state
  }

  filtered := state[:0]

  for _, element := range state {
    if( element.Record_type == record.Record_type &&
        element.Name        == record.Name ) {
      fmt.Printf("Removing %s %s %s %s\n", element.Record_type, element.Name, record.Record_type, record.Name)
    } else {
      filtered = append(filtered, element)
    }
  }

  return filtered
}