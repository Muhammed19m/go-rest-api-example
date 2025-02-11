package model

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tansaction := Transaction{1, DEPOSIT, 10}
	if err:= tansaction.Validate(); err != nil {
		t.Error("unexpected error, found ", err)
	}

	tansaction = Transaction{0, DEPOSIT, 10}
	if err:= tansaction.Validate(); !errors.Is(err, ErrInvalidId) {
		t.Error("expected error ErrInvalidId, found ", err)
	} 

	tansaction = Transaction{-1, DEPOSIT, 10}
	if err:= tansaction.Validate(); !errors.Is(err, ErrInvalidId) {
		t.Error("expected error ErrInvalidId, found ", err)
	}

	tansaction = Transaction{1, DEPOSIT, 0}
	if err:= tansaction.Validate(); err != nil {
		t.Error("unexpected error, found ", err)
	}

	tansaction = Transaction{1, DEPOSIT, -1}
	if err:= tansaction.Validate(); !errors.Is(err, ErrInvalidAmount) {
			t.Error("expected error ErrInvalidAmount, found ", err)		
	}
	
	tansaction = Transaction{1, "deposit", 1}
	if err:= tansaction.Validate(); !errors.Is(err, ErrInvalidOperationType) {
			t.Error("expected error ErrInvalidOperationType, found ", err)		
	}

}


func TestValidateUUid(t *testing.T) {
	tests := []struct{
				id int
				expectedError error
			} {
		{1, nil},
		{100, nil},
		{0, ErrInvalidId},
		{-1, ErrInvalidId},
		{-100, ErrInvalidId},
	}
	for _, tt := range tests {
		if err := ValidateUUid(tt.id); err != tt.expectedError {

			t.Error("expected ", tt.expectedError, " found ", err)
		}
	}
}