package utils

import (
  "encoding/base64"
  "github.com/satori/go.uuid"
  )


  func Encode(id *uuid.UUID) string {
    return base64.RawURLEncoding.EncodeToString(id.Bytes())
  }

  func Decode(id string) (*uuid.UUID, error) {
    dec, err := base64.RawURLEncoding.DecodeString(id)
    if err != nil {
      return nil, err
    }
    decID, err := uuid.FromBytes(dec)
    if err != nil {
      return nil, err
    }
    return &decID, nil
  }