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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.fields.wb {

				wr, err := tt.fields.b.NewWriter()
				if (err != nil) != tt.wantErr {
					t.Errorf("Batch.NewWriter() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				go (func() {
					wr.Write(tt.fields.wb[i])
					defer wr.Close()
				})()
			}

			got := []byte{}
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
