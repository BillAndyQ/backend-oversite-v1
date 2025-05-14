package domain

type OrdenTrabajoPersona struct {
	ID                     uint   `json:"id" gorm:"primaryKey"`
	OrdenTrabajoID         string `json:"orden_trabajo_id"`
	Empresa                string `json:"empresa"`
	EmpresaMatriz          string `json:"empresa_matriz"`
	Ruc                    string `json:"ruc"`
	FechaServicio          string `json:"fecha_servicio"`
	Certificadora          string `json:"certificadora"`
	TipoCurso              string `json:"tipo_curso"`
	Modalidad              string `json:"modalidad"`
	NombreCurso            string `json:"nombre_curso"`
	NombreInstructor       string `json:"nombre_instructor"`
	Fotocheck              string `json:"fotocheck"`
	DNIalumno              string `json:"dni_alumno"`
	NombresApellidosAlumno string `json:"nombres_apellidos_alumno"`

	AproboCurso bool   `json:"aprobo_curso"`
	Nota        string `json:"nota"`
	Puesto      string `json:"puesto"`

	EquipoDeExamen string `json:"equipo_de_examen"`

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
