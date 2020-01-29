package main

func parseSignaturesFile(path string) (map[string]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  sigs := make(map[string]string)
  scanner := bufio.NewScanner(file)
  for lnum := 1; scanner.Scan(); lnum++
}
