package main

func parseSignaturesFile(path string) (map[string]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  sigs := make(map[string]string)
  scanner := bufio.NewScanner(file)
  for lnum := 1; scanner.Scan(); lnum++　{
    fields := strings.Fields(scanner.Text())
    if len(fields) != 2 {
      return nil, fmt.Errorf("%s: %d bad line\n", path, lnum)
    }
    sigs[fields[1]] = fields[0]
  }

  if err := scanner.Err(); err != nil {
    return nil, err
  }
  return sigs, nil
  }
}

func fileMD5(path string) (string, error) {
  file, err := os.Open(path)
  if err != nil {
    return "", err
  }
  defer file.Close()

  hash := md5.New()
  if _, err := io.Copy(hash, file); err != nil {
    return "", err
  }
  return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

type result struct {
  path string
  match bool
  err error
}
