package kafka

type Message struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type   string `json:"type"`
			Fields []struct {
				Type       string `json:"type"`
				Optional   bool   `json:"optional"`
				Field      string `json:"field"`
				Name       string `json:"name,omitempty"`
				Version    int    `json:"version,omitempty"`
				Parameters struct {
					Allowed string `json:"allowed"`
				} `json:"parameters,omitempty"`
				Default string `json:"default,omitempty"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field"`
			Version  int    `json:"version,omitempty"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
		Version  int    `json:"version"`
	} `json:"schema"`
	Payload struct {
		Before interface{} `json:"before"`
		After  interface{} `json:"after"`
		Source struct {
			Version   string      `json:"version"`
			Connector string      `json:"connector"`
			Name      string      `json:"name"`
			TsMs      int64       `json:"ts_ms"`
			Snapshot  string      `json:"snapshot"`
			Db        string      `json:"db"`
			Sequence  string      `json:"sequence"`
			Schema    string      `json:"schema"`
			Table     string      `json:"table"`
			TxId      int         `json:"txId"`
			Lsn       int         `json:"lsn"`
			Xmin      interface{} `json:"xmin"`
		} `json:"source"`
		Op          string      `json:"op"`
		TsMs        int64       `json:"ts_ms"`
		Transaction interface{} `json:"transaction"`
	} `json:"payload"`
}

type RedditCommentMessage struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type   string `json:"type"`
			Fields []struct {
				Type       string `json:"type"`
				Optional   bool   `json:"optional"`
				Field      string `json:"field"`
				Name       string `json:"name,omitempty"`
				Version    int    `json:"version,omitempty"`
				Parameters struct {
					Allowed string `json:"allowed"`
				} `json:"parameters,omitempty"`
				Default string `json:"default,omitempty"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field"`
			Version  int    `json:"version,omitempty"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
		Version  int    `json:"version"`
	} `json:"schema"`
	Payload struct {
		Before interface{}   `json:"before"`
		After  RedditComment `json:"after"`
		Source struct {
			Version   string      `json:"version"`
			Connector string      `json:"connector"`
			Name      string      `json:"name"`
			TsMs      int64       `json:"ts_ms"`
			Snapshot  string      `json:"snapshot"`
			Db        string      `json:"db"`
			Sequence  string      `json:"sequence"`
			Schema    string      `json:"schema"`
			Table     string      `json:"table"`
			TxId      int         `json:"txId"`
			Lsn       int         `json:"lsn"`
			Xmin      interface{} `json:"xmin"`
		} `json:"source"`
		Op          string      `json:"op"`
		TsMs        int64       `json:"ts_ms"`
		Transaction interface{} `json:"transaction"`
	} `json:"payload"`
}

type RedditSubmissionMessage struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type   string `json:"type"`
			Fields []struct {
				Type       string `json:"type"`
				Optional   bool   `json:"optional"`
				Field      string `json:"field"`
				Name       string `json:"name,omitempty"`
				Version    int    `json:"version,omitempty"`
				Parameters struct {
					Allowed string `json:"allowed"`
				} `json:"parameters,omitempty"`
				Default string `json:"default,omitempty"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field"`
			Version  int    `json:"version,omitempty"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
		Version  int    `json:"version"`
	} `json:"schema"`
	Payload struct {
		Before interface{}      `json:"before"`
		After  RedditSubmission `json:"after"`
		Source struct {
			Version   string      `json:"version"`
			Connector string      `json:"connector"`
			Name      string      `json:"name"`
			TsMs      int64       `json:"ts_ms"`
			Snapshot  string      `json:"snapshot"`
			Db        string      `json:"db"`
			Sequence  string      `json:"sequence"`
			Schema    string      `json:"schema"`
			Table     string      `json:"table"`
			TxId      int         `json:"txId"`
			Lsn       int         `json:"lsn"`
			Xmin      interface{} `json:"xmin"`
		} `json:"source"`
		Op          string      `json:"op"`
		TsMs        int64       `json:"ts_ms"`
		Transaction interface{} `json:"transaction"`
	} `json:"payload"`
}
