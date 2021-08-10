package bibatch

import (
	"testing"
)

func TestBatch_NewWriter(t *testing.T) {
	type fields struct {
		b  *Batch
		wb [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Single item",
			fields: fields{
				b: NewBatch(1),
				wb: [][]byte{
					[]byte("Hello Batch"),
				},
			},
			want:    []byte("Hello Batch"),
			wantErr: false,
		},
		{
			name: "few item",
			fields: fields{
				b: NewBatch(2),
				wb: [][]byte{
					[]byte("Hello"),
					[]byte(" Batch"),
				},
			},
			want:    []byte("Hello Batch"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l int
			for i := range tt.fields.wb {

				wr, err := tt.fields.b.NewWriter()
				if (err != nil) != tt.wantErr {
					t.Errorf("Batch.NewWriter() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				l += len(tt.fields.wb[i])

				go (func(buff []byte) {
					wr.Write(buff)
					defer wr.Close()
				})(tt.fields.wb[i])
			}

			got := make([]byte, l)
			n, err := tt.fields.b.Read(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Batch.NewWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("Batch.NewWriter() = %v(%v), want %v", got, n, tt.want)
			}
		})
	}
}
