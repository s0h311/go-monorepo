package utils

func TryOrThrow(callback func() error) {
  err := callback()

  if err != nil {
    panic(err)
  }
}
