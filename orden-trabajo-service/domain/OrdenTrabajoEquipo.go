package domain

type OrdenTrabajoEquipo struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	OrdenTrabajoID    string `json:"orden_trabajo_id"`
	OrdenServiceID    string `json:"orden_servicio_id"`
	Empresa           string `json:"empresa"`
	EmpresaMatriz     string `json:"empresa_matriz"`
	Ruc               string `json:"ruc"`
	FechaServicio     string `json:"fecha_servicio"`
	Certificadora     string `json:"certificadora"`
	TipoUnidad        string `json:"tipo_unidad"`
	Placa             string `json:"placa"`
	Area              string `json:"area"`
	NFactura          string `json:"n_factura"`
	FechaFactura      string `json:"fecha_factura"`
	MontoSinIGV       string `json:"monto_sin_igv"`
	MontoConIGV       string `json:"monto_con_igv"`
	IGV               string `json:"igv"`
	MontoDetraccion   string `json:"detraccion"`
	EnSoles           string `json:"en_soles"`
	EnDolares         string `json:"en_dolares"`
	Moneda            string `json:"moneda"`
	Estado            string `json:"estado"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_modificacion"`
}
