// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Annotation) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Annotation) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Time.Set {
			e.FieldStart("time")
			s.Time.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.EndTime.Set {
			e.FieldStart("endTime")
			s.EndTime.Encode(e, json.EncodeDateTime)
		}
	}
	{
		e.FieldStart("type")
		e.Str(s.Type)
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		e.FieldStart("datasets")
		e.ArrStart()
		for _, elem := range s.Datasets {
			e.Str(elem)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
}

var jsonFieldsNameOfAnnotation = [8]string{
	0: "title",
	1: "description",
	2: "time",
	3: "endTime",
	4: "type",
	5: "url",
	6: "datasets",
	7: "id",
}

// Decode decodes Annotation from json.
func (s *Annotation) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Annotation to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "time":
			if err := func() error {
				s.Time.Reset()
				if err := s.Time.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		case "endTime":
			if err := func() error {
				s.EndTime.Reset()
				if err := s.EndTime.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"endTime\"")
			}
		case "type":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Type = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "datasets":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				s.Datasets = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Datasets = append(s.Datasets, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"datasets\"")
			}
		case "id":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Annotation")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b11010000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAnnotation) {
					name = jsonFieldsNameOfAnnotation[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Annotation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Annotation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAnnotationsOKApplicationJSON as json.
func (s GetAnnotationsOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []Annotation(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes GetAnnotationsOKApplicationJSON from json.
func (s *GetAnnotationsOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAnnotationsOKApplicationJSON to nil")
	}
	var unwrapped []Annotation
	if err := func() error {
		unwrapped = make([]Annotation, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Annotation
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAnnotationsOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s GetAnnotationsOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAnnotationsOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *NewAnnotation) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *NewAnnotation) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Time.Set {
			e.FieldStart("time")
			s.Time.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.EndTime.Set {
			e.FieldStart("endTime")
			s.EndTime.Encode(e, json.EncodeDateTime)
		}
	}
	{
		e.FieldStart("type")
		e.Str(s.Type)
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		e.FieldStart("datasets")
		e.ArrStart()
		for _, elem := range s.Datasets {
			e.Str(elem)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfNewAnnotation = [7]string{
	0: "title",
	1: "description",
	2: "time",
	3: "endTime",
	4: "type",
	5: "url",
	6: "datasets",
}

// Decode decodes NewAnnotation from json.
func (s *NewAnnotation) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode NewAnnotation to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "time":
			if err := func() error {
				s.Time.Reset()
				if err := s.Time.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		case "endTime":
			if err := func() error {
				s.EndTime.Reset()
				if err := s.EndTime.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"endTime\"")
			}
		case "type":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Type = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "datasets":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				s.Datasets = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Datasets = append(s.Datasets, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"datasets\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode NewAnnotation")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b01010000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfNewAnnotation) {
					name = jsonFieldsNameOfNewAnnotation[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *NewAnnotation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *NewAnnotation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UpdateAnnotation) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UpdateAnnotation) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Time.Set {
			e.FieldStart("time")
			s.Time.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.EndTime.Set {
			e.FieldStart("endTime")
			s.EndTime.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		if s.Datasets != nil {
			e.FieldStart("datasets")
			e.ArrStart()
			for _, elem := range s.Datasets {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfUpdateAnnotation = [7]string{
	0: "title",
	1: "description",
	2: "time",
	3: "endTime",
	4: "type",
	5: "url",
	6: "datasets",
}

// Decode decodes UpdateAnnotation from json.
func (s *UpdateAnnotation) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateAnnotation to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "time":
			if err := func() error {
				s.Time.Reset()
				if err := s.Time.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		case "endTime":
			if err := func() error {
				s.EndTime.Reset()
				if err := s.EndTime.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"endTime\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "datasets":
			if err := func() error {
				s.Datasets = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Datasets = append(s.Datasets, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"datasets\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateAnnotation")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateAnnotation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateAnnotation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
