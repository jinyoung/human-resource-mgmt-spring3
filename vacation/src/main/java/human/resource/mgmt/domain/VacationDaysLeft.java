package human.resource.mgmt.domain;

import human.resource.mgmt.VacationApplication;
import human.resource.mgmt.domain.VacationDaysInsufficient;
import human.resource.mgmt.domain.VacationDaysIntialized;
import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;

@Entity
@Table(name = "VacationDaysLeft_table")
@Data
public class VacationDaysLeft {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private String userId;

    private Integer dayCount;

    @PostPersist
    public void onPostPersist() {
        VacationDaysIntialized vacationDaysIntialized = new VacationDaysIntialized(
            this
        );
        vacationDaysIntialized.publishAfterCommit();

        VacationDaysInsufficient vacationDaysInsufficient = new VacationDaysInsufficient(
            this
        );
        vacationDaysInsufficient.publishAfterCommit();
    }

    public static VacationDaysLeftRepository repository() {
        VacationDaysLeftRepository vacationDaysLeftRepository = VacationApplication.applicationContext.getBean(
            VacationDaysLeftRepository.class
        );
        return vacationDaysLeftRepository;
    }

    public void add(AddCommand addCommand) {
        VacationDaysAdded vacationDaysAdded = new VacationDaysAdded(this);
        vacationDaysAdded.publishAfterCommit();
    }

    public void use(UseCommand useCommand) {
        VacationDaysUsed vacationDaysUsed = new VacationDaysUsed(this);
        vacationDaysUsed.publishAfterCommit();
    }
}
