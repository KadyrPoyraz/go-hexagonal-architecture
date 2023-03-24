package api

import (
    "github.com/KadyrPoyraz/go-hexagonal-architecture/internal/ports"
)

type Adapter struct {
    db    ports.DbPort
    arith ports.ArithmeticPort
}

func newAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
    return &Adapter{
        db: db,
        arith: arith,
    }
}

func (apia Adapter) getAddition(a, b int32) (int32, error) {
    result, err := apia.arith.Addition(a, b)
    if err != nil {
        return 0, err
    }

    err = apia.db.AddToHistory(result, "addition")
    if err != nil {
        return 0, err
    }

    return result, err
}

func (apia Adapter) getSubstruction(a, b int32) (int32, error) {
    result, err := apia.arith.Subtraction(a, b)
    if err != nil {
        return 0, err
    }

    err = apia.db.AddToHistory(result, "substruction")
    if err != nil {
        return 0, err
    }

    return result, err
}

func (apia Adapter) getDivision(a, b int32) (int32, error) {
    result, err := apia.arith.Division(a, b)
    if err != nil {
        return 0, err
    }

    err = apia.db.AddToHistory(result, "division")
    if err != nil {
        return 0, err
    }

    return result, err
}

func (apia Adapter) getMultiplication(a, b int32) (int32, error) {
    result, err := apia.arith.Multiplication(a, b)
    if err != nil {
        return 0, err
    }

    err = apia.db.AddToHistory(result, "multiplication")
    if err != nil {
        return 0, err
    }

    return result, err
}
