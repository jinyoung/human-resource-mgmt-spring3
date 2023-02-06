package human.resource.mgmt.domain;

import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;

@Entity
@Table(name = "VacationDaysStatus_table")
@Data
public class VacationDaysStatus {

    @Id
    //@GeneratedValue(strategy=GenerationType.AUTO)
    private String userId;

    private Integer daysLeft;
}
